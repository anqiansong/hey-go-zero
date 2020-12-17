// Code generated by goctl. DO NOT EDIT!
// Source: course.proto

//go:generate mockgen -destination ./courseservice_mock.go -package courseservice -source $GOFILE

package courseservice

import (
	"context"

	"hey-go-zero/service/course/rpc/course"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	Course          = course.Course
	CourseListReply = course.CourseListReply
	IdReq           = course.IdReq
	IdsReq          = course.IdsReq

	CourseService interface {
		//  查询课程
		FindOne(ctx context.Context, in *IdReq) (*Course, error)
		//  批量获取课程
		FindByIds(ctx context.Context, in *IdsReq) (*CourseListReply, error)
	}

	defaultCourseService struct {
		cli zrpc.Client
	}
)

func NewCourseService(cli zrpc.Client) CourseService {
	return &defaultCourseService{
		cli: cli,
	}
}

//  查询课程
func (m *defaultCourseService) FindOne(ctx context.Context, in *IdReq) (*Course, error) {
	client := course.NewCourseServiceClient(m.cli.Conn())
	return client.FindOne(ctx, in)
}

//  批量获取课程
func (m *defaultCourseService) FindByIds(ctx context.Context, in *IdsReq) (*CourseListReply, error) {
	client := course.NewCourseServiceClient(m.cli.Conn())
	return client.FindByIds(ctx, in)
}
