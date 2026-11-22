package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type TlpassStackProps struct {
	awscdk.StackProps
}

func NewTlpassStack(scope constructs.Construct, id string, props *TlpassStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	return awscdk.NewStack(scope, &id, &sprops)
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewTlpassStack(app, "tlpass-dev", &TlpassStackProps{
		StackProps: awscdk.StackProps{
			Env: &awscdk.Environment{
				Account: jsii.String("973679386229"),
				Region:  jsii.String("eu-central-1"),
			},
			Tags: &map[string]*string{
				"app": jsii.String("tlpass"),
				"env": jsii.String("dev"),
			},
		},
	})

	app.Synth(nil)
}