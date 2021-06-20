package delivery

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"linkShortener/internal/pkg/shortLink"
	pb "linkShortener/proto"
)

type ShortLinkServer struct {
	Uc shortLink.ShortLinkUsecase
}

func (sls *ShortLinkServer) Create(ctx context.Context, ol *pb.OriginalLink) (*pb.ShortenLink, error) {
	shortenLink, err := sls.Uc.Create(ol.OriginalLink)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ShortenLink{
		ShortenLink: shortenLink,
	}, nil
}

func (sls *ShortLinkServer) Get(ctx context.Context, sl *pb.ShortenLink) (*pb.OriginalLink, error) {
	originalLink, err := sls.Uc.Get(sl.ShortenLink)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &pb.OriginalLink{
		OriginalLink: originalLink,
	}, nil
}



