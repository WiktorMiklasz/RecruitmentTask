package test

import (
  	"testing"

  	"github.com/gruntwork-io/terratest/modules/docker"
  	"github.com/stretchr/testify/assert"
)

func TestNginx(t *testing.T){
	tag := "testimage"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}
	docker.Build(t, "../angproject/", buildOptions)

	opts := &docker.RunOptions{Command: []string{"nginx", "-v"}}
	output := docker.Run(t, tag, opts)
	assert.Contains(t, output ,"nginx version: nginx/1.21.6")
}
