package impl

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/morphcloud/advert-service/internal/pb"
)

var adverts = []*pb.Advert{
	{
		Id:          1,
		CategoryId:  1,
		UserId:      1,
		Title:       "TDD",
		Description: "WTF",
		Price:       500,
		Status:      1,
		ActivatedAt: getTimestamp(),
		CreatedAt:   getTimestamp(),
		UpdatedAt:   getTimestamp(),
	},
	{
		Id:          2,
		CategoryId:  2,
		UserId:      2,
		Title:       "QKA",
		Description: "AERJM",
		Price:       5000,
		Status:      1,
		ActivatedAt: getTimestamp(),
		CreatedAt:   getTimestamp(),
		UpdatedAt:   getTimestamp(),
	},
	{
		Id:          3,
		CategoryId:  1,
		UserId:      2,
		Title:       "EAEM",
		Description: "MFAJER",
		Price:       3300,
		Status:      2,
		ActivatedAt: getTimestamp(),
		CreatedAt:   getTimestamp(),
		UpdatedAt:   getTimestamp(),
	},
}

type advertServiceServer struct{}

func NewAdvertServiceServer() *advertServiceServer {
	return &advertServiceServer{}
}

func (a *advertServiceServer) Create(ctx context.Context, advert *pb.Advert) (*pb.Advert, error) {
	return adverts[0], nil
}

func (a *advertServiceServer) Read(ctx context.Context, filter *pb.Filter) (*pb.Advert, error) {
	return adverts[filter.Id-1], nil
}

func (a *advertServiceServer) ReadAll(empty *empty.Empty, stream pb.AdvertService_ReadAllServer) error {
	for _, advert := range adverts {
		if err := stream.Send(advert); err != nil {
			return nil
		}
	}
	return nil
}

func (a *advertServiceServer) Update(ctx context.Context, advert *pb.Advert) (*pb.Advert, error) {
	adv := adverts[advert.Id-1]

	if advert.CategoryId != 0 {
		adv.CategoryId = advert.CategoryId
	}
	if advert.Status != 0 {
		adv.Status = advert.Status
	}
	if advert.Price != 0 {
		adv.Price = advert.Price
	}
	if advert.Title != "" {
		adv.Title = advert.Title
	}
	if advert.Description != "" {
		adv.Description = advert.Description
	}

	return adverts[advert.Id-1], nil
}

func (a *advertServiceServer) Delete(ctx context.Context, advert *pb.Advert) (*pb.Advert, error) {
	return adverts[advert.Id-1], nil
}

func getTimestamp() *timestamp.Timestamp {
	tmstmp, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return tmstmp
}
