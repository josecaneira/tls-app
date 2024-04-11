# Copyright 2022 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
#
# Original work by Google Inc. at https://github.com/googlecloudplatform/kubernetes-engine-samples/tree/394477ca2a99/quickstarts/hello-app-tls
# Derivate work by José Caneira at https://github.com/josecaneira for TLS end-to-end Demo.
#

FROM golang:1.21.0 as builder
WORKDIR /app
RUN go mod init hello-app-tls
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /hello-app-tls

#FROM gcr.io/distroless/base-debian11
FROM mcr.microsoft.com/devcontainers/base:bullseye
WORKDIR /
COPY --from=builder /hello-app-tls /hello-app-tls
#USER nonroot:nonroot
USER nobody:nobody
CMD ["/hello-app-tls"]
