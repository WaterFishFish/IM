package logic

import (
	"context"
	"easy-chat/apps/user/rpc/user"
	"reflect"
	"testing"
)

func TestRegisterLogic_Register(t *testing.T) {

	type args struct {
		in *user.RegisterReq
	}
	tests := []struct {
		name      string
		args      args
		wantPrint bool
		wantErr   bool
	}{
		{
			name: "1", args: args{in: &user.RegisterReq{
				Phone:    "123456789",
				Nickname: "FISHWATER",
				Password: "123456",
				Avatar:   "png.jpg",
				Sex:      0,
			}}, wantPrint: true, wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewRegisterLogic(context.Background(), svcCtx)
			got, err := l.Register(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantPrint) {
				t.Log(tt.name, got)
			}
		})
	}
}
