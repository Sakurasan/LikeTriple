package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func Test_getCoinVideo(t *testing.T) {
	type args struct {
		vmid string
	}
	tests := []struct {
		name    string
		args    args
		wantCv  *coinVideo
		wantErr bool
	}{
		{
			name: "coinVideo",
			args: args{vmid: "208259"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCv, err := getCoinVideo(tt.args.vmid)
			if err != nil {
				t.Errorf("getCoinVideo() error = %v", err)
				return
			}
			for _, v := range gotCv.Data {
				fmt.Println(v.Title, v.ShortLink, v.Aid)
			}
		})
	}
}

func Test_filelist(t *testing.T) {
	fileinfo, err := ioutil.ReadDir("/Users/nh/Code/Go/src/LikeTriple")
	if err != nil {
		t.Errorf("fileinfo error = %v", err)
	}
	for _, f := range fileinfo {
		fmt.Println(f.Name())
	}
}

func Test_fileISexist(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "file exist",
			args: args{name: "README.md"},
			want: true,
		},
	}
	str, _ := os.Getwd()
	fmt.Println(str)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileISexist(tt.args.name); got != tt.want {
				t.Errorf("fileISexist() = %v, want %v", got, tt.want)
			}
		})
	}
}
