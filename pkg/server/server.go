// Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

// Serve start the webserver and return a <context.cancelFunc> to gracefully shutdown the server.
func Serve(bindAddress string, port int, logger *logrus.Logger) context.CancelFunc {
	http.Handle("/metrics", promhttp.Handler())
	server := http.Server{
		Addr: fmt.Sprintf("%s:%d", bindAddress, port),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	go server.ListenAndServe()
	logger.Infof("webserver is running on port %d", port)
	server.Shutdown(ctx)
	return cancel
}