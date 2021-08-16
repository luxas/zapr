/*
Copyright 2019 The logr Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

type StringErr string

func (e StringErr) Error() string { return string(e) }

func Helper(log logr.Logger, msg string) {
	helper2(log, msg)
}

func helper2(log logr.Logger, msg string) {
	logr.WithCallDepth(log, 2).Info(msg)
}

func main() {
	log := zapr.NewLogger(zap.NewExample())
	log = log.WithName("MyName").WithValues("user", "you")
	log.Info("hello", "val1", 1, "val2", map[string]int{"k": 1})
	log.V(1).Info("you should see this")
	log.V(1).V(1).Info("you should NOT see this")
	log.V(1).WithName("Verbose").V(1).Info("you should NOT see this")
	log.V(1).WithValues("foo", true).V(1).Info("you should NOT see this")
	Helper(log.V(2), "you should NOT see this")
	log.Error(nil, "uh oh", "trouble", true, "reasons", []float64{0.1, 0.11, 3.14})
	log.Error(StringErr("an error occurred"), "goodbye", "code", -1)
	Helper(log, "thru a helper")
}

/*
	Expected output:

	{"level":"info","logger":"MyName","msg":"hello","user":"you","val1":1,"val2":{"k":1}}
	{"level":"debug","logger":"MyName","msg":"you should see this","user":"you"}
	{"level":"error","logger":"MyName","msg":"uh oh","user":"you","trouble":true,"reasons":[0.1,0.11,3.14]}
	{"level":"error","logger":"MyName","msg":"goodbye","user":"you","code":-1,"error":"an error occurred"}
	{"level":"info","logger":"MyName","msg":"thru a helper","user":"you"}
*/
