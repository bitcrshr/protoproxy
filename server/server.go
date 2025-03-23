package server

import (
	_ "embed"
	"net"

	"github.com/bitcrshr/protoproxy/pathmap"
	"github.com/rotisserie/eris"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	//go:embed desc.binpb
	rawDescriptorSet []byte

	descriptorSet = &descriptorpb.FileDescriptorSet{}

	pathMap pathmap.PathMap
)

func Serve(port string) {
	startup()

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		err = eris.Wrapf(err, "failed to start listener on port %s", port)
		log.Panic().Err(err).Msg("Serve failed")
	}

	server := grpc.NewServer(
		grpc.UnknownServiceHandler(handle),
	)

	log.Info().Msgf("server running on port %s :>\n", port)
	if err := server.Serve(lis); err != nil {
		err = eris.Wrapf(err, "failed to serve with listener for port %s", port)
		log.Panic().Err(err).Msg("Serve failed")
	}
}

func handle(srv any, stream grpc.ServerStream) error {
	method, ok := grpc.MethodFromServerStream(stream)
	if !ok || method == "" {
		return notFound(method, stream)
	}

	log.Info().Msgf("handling request: %s\n", method)

	reqJson, err := pathMap.ParseRequest(method)
	if err != nil {
		log.Error().Err(err).Msg("handle failed")
		return err
	}

	log.Debug().RawJSON("request_json", reqJson).Msg("handle: got msg json")

	res := &emptypb.Empty{}
	if err := stream.SendMsg(res); err != nil {
		err = eris.Wrap(err, "failed to send response")
		log.Error().Err(err).Any("res", res).Msg("handle failed")
		return err
	}

	return nil
}

func notFound(method string, stream grpc.ServerStream) error {
	return stream.SendMsg(
		status.Errorf(codes.NotFound, "grpc method %s not found", method),
	)
}

func startup() {
	if err := proto.Unmarshal(rawDescriptorSet, descriptorSet); err != nil {
		err = eris.Wrap(err, "failed to unmarshal binary file descriptor set")
		log.Panic().Err(err).Msg("server startup failed")
	}

	var err error
	if pathMap, err = pathmap.New(descriptorSet); err != nil {
		err = eris.Wrap(err, "failed to construct pathmap")
		log.Panic().Err(err).Msg("server startup failed")
	}

	log.Debug().Any("pathMap", pathMap).Send()
}
