Last login: Sat Mar 20 23:20:52 on ttys001

The default interactive shell is now zsh.
To update your account to use zsh, please run `chsh -s /bin/zsh`.
For more details, please visit https://support.apple.com/kb/HT208050.
LLUE-M-C5ME:~ llue$ protoc
Usage: protoc [OPTION] PROTO_FILES
Parse PROTO_FILES and generate output based on the options given:
  -IPATH, --proto_path=PATH   Specify the directory in which to search for
                              imports.  May be specified multiple times;
                              directories will be searched in order.  If not
                              given, the current working directory is used.
                              If not found in any of the these directories,
                              the --descriptor_set_in descriptors will be
                              checked for required proto file.
  --version                   Show version info and exit.
  -h, --help                  Show this text and exit.
  --encode=MESSAGE_TYPE       Read a text-format message of the given type
                              from standard input and write it in binary
                              to standard output.  The message type must
                              be defined in PROTO_FILES or their imports.
  --decode=MESSAGE_TYPE       Read a binary message of the given type from
                              standard input and write it in text format
                              to standard output.  The message type must
                              be defined in PROTO_FILES or their imports.
  --decode_raw                Read an arbitrary protocol message from
                              standard input and write the raw tag/value
                              pairs in text format to standard output.  No
                              PROTO_FILES should be given when using this
                              flag.
  --descriptor_set_in=FILES   Specifies a delimited list of FILES
                              each containing a FileDescriptorSet (a
                              protocol buffer defined in descriptor.proto).
                              The FileDescriptor for each of the PROTO_FILES
                              provided will be loaded from these
                              FileDescriptorSets. If a FileDescriptor
                              appears multiple times, the first occurrence
                              will be used.
  -oFILE,                     Writes a FileDescriptorSet (a protocol buffer,
    --descriptor_set_out=FILE defined in descriptor.proto) containing all of
                              the input files to FILE.
  --include_imports           When using --descriptor_set_out, also include
                              all dependencies of the input files in the
                              set, so that the set is self-contained.
  --include_source_info       When using --descriptor_set_out, do not strip
                              SourceCodeInfo from the FileDescriptorProto.
                              This results in vastly larger descriptors that
                              include information about the original
                              location of each decl in the source file as
                              well as surrounding comments.
  --dependency_out=FILE       Write a dependency output file in the format
                              expected by make. This writes the transitive
                              set of input file paths to FILE
  --error_format=FORMAT       Set the format in which to print errors.
                              FORMAT may be 'gcc' (the default) or 'msvs'
                              (Microsoft Visual Studio format).
  --print_free_field_numbers  Print the free field numbers of the messages
                              defined in the given proto files. Groups share
                              the same field number space with the parent 
                              message. Extension ranges are counted as 
                              occupied fields numbers.

  --plugin=EXECUTABLE         Specifies a plugin executable to use.
                              Normally, protoc searches the PATH for
                              plugins, but you may specify additional
                              executables not in the path using this flag.
                              Additionally, EXECUTABLE may be of the form
                              NAME=PATH, in which case the given plugin name
                              is mapped to the given executable even if
                              the executable's own name differs.
  --cpp_out=OUT_DIR           Generate C++ header and source.
  --csharp_out=OUT_DIR        Generate C# source file.
  --java_out=OUT_DIR          Generate Java source file.
  --js_out=OUT_DIR            Generate JavaScript source.
  --objc_out=OUT_DIR          Generate Objective C header and source.
  --php_out=OUT_DIR           Generate PHP source file.
  --python_out=OUT_DIR        Generate Python source file.
  --ruby_out=OUT_DIR          Generate Ruby source file.
  @<filename>                 Read options and filenames from file. If a
                              relative file path is specified, the file
                              will be searched in the working directory.
                              The --proto_path option will not affect how
                              this argument file is searched. Content of
                              the file will be expanded in the position of
                              @<filename> as in the argument list. Note
                              that shell expansion is not applied to the
                              content of the file (i.e., you cannot use
                              quotes, wildcards, escapes, commands, etc.).
                              Each line corresponds to a single argument,
                              even if it contains spaces.
LLUE-M-C5ME:~ llue$ ls
10.24.81.151				contact
2020-06-11.saz				docker
2020-07-07.saz				evtlog.pb.go
Applications				fex_sensor.bin
Calibre Library				go
Desktop					goLang
Documents				grpc-go
Downloads				iCloud Drive (Archive)
Google Drive				iCloud Drive (Archive) - 1
Header					kafka_2.12-2.4.0
Library					mac_address
Message					macbook_list
Movies					mei_card
Music					notes
MyJabberFiles				nxos.7.0.3.IHD8.0.703.bin
Pictures				pack
Postman					protobuf-3.15.1
Public					python_test
Sites					recovery_key
Walkme					redis
a					room
aa					stop_crash
bazel-3.7.0-installer-darwin-x86_64.sh	telemetry
bb					test
bin					tmp
cc					unit-test-3
LLUE-M-C5ME:~ llue$ mkdir evtlog
LLUE-M-C5ME:~ llue$ vim evtlog.proto
LLUE-M-C5ME:~ llue$ protoc -go_out=. *.proto
Unknown flag: -g
LLUE-M-C5ME:~ llue$ protoc --go_out=. *.proto
protoc-gen-go: program not found or is not executable
Please specify a program using absolute path or make sure the program is available in your PATH system variable
--go_out: protoc-gen-go: Plugin failed with status code 1.
LLUE-M-C5ME:~ llue$ which protoc
/usr/local/bin/protoc
LLUE-M-C5ME:~ llue$ cd /usr/local/go
LLUE-M-C5ME:go llue$ ls
AUTHORS		LICENSE		SECURITY.md	bin		lib		robots.txt
CONTRIBUTING.md	PATENTS		VERSION		doc		misc		src
CONTRIBUTORS	README.md	api		favicon.ico	pkg		test
LLUE-M-C5ME:go llue$ cd bin
LLUE-M-C5ME:bin llue$ ls
go	gofmt
LLUE-M-C5ME:bin llue$ cd go
-bash: cd: go: Not a directory
LLUE-M-C5ME:bin llue$ ls
go	gofmt
LLUE-M-C5ME:bin llue$ cd go
-bash: cd: go: Not a directory
LLUE-M-C5ME:bin llue$ cd 
LLUE-M-C5ME:~ llue$ cd go_grpc
-bash: cd: go_grpc: No such file or directory
LLUE-M-C5ME:~ llue$ cd grpc_go
-bash: cd: grpc_go: No such file or directory
LLUE-M-C5ME:~ llue$ ls
10.24.81.151				docker
2020-06-11.saz				evtlog
2020-07-07.saz				evtlog.pb.go
Applications				evtlog.proto
Calibre Library				fex_sensor.bin
Desktop					go
Documents				goLang
Downloads				grpc-go
Google Drive				iCloud Drive (Archive)
Header					iCloud Drive (Archive) - 1
Library					kafka_2.12-2.4.0
Message					mac_address
Movies					macbook_list
Music					mei_card
MyJabberFiles				notes
Pictures				nxos.7.0.3.IHD8.0.703.bin
Postman					pack
Public					protobuf-3.15.1
Sites					python_test
Walkme					recovery_key
a					redis
aa					room
bazel-3.7.0-installer-darwin-x86_64.sh	stop_crash
bb					telemetry
bin					test
cc					tmp
contact					unit-test-3
LLUE-M-C5ME:~ llue$ cd grpc-go
LLUE-M-C5ME:grpc-go llue$ ls
AUTHORS					health
CODE-OF-CONDUCT.md			install_gae.sh
CONTRIBUTING.md				interceptor.go
Documentation				internal
GOVERNANCE.md				interop
LICENSE					keepalive
MAINTAINERS.md				metadata
Makefile				peer
README.md				picker_wrapper.go
SECURITY.md				picker_wrapper_test.go
attributes				pickfirst.go
backoff					pickfirst_test.go
backoff.go				preloader.go
balancer				profiling
balancer_conn_wrappers.go		reflection
balancer_conn_wrappers_test.go		regenerate.sh
balancer_switching_test.go		resolver
benchmark				resolver_conn_wrapper.go
binarylog				resolver_conn_wrapper_test.go
call.go					rpc_util.go
call_test.go				rpc_util_test.go
channelz				security
clientconn.go				server.go
clientconn_state_transition_test.go	server_test.go
clientconn_test.go			service_config.go
cmd					service_config_test.go
codec.go				serviceconfig
codec_test.go				stats
codegen.sh				status
codes					stream.go
connectivity				stress
credentials				tap
dialoptions.go				test
doc.go					testdata
encoding				trace.go
examples				trace_test.go
go.mod					version.go
go.sum					vet.sh
grpc_test.go				xds
grpclog
LLUE-M-C5ME:grpc-go llue$ cd
LLUE-M-C5ME:~ llue$ ls
10.24.81.151				docker
2020-06-11.saz				evtlog
2020-07-07.saz				evtlog.pb.go
Applications				evtlog.proto
Calibre Library				fex_sensor.bin
Desktop					go
Documents				goLang
Downloads				grpc-go
Google Drive				iCloud Drive (Archive)
Header					iCloud Drive (Archive) - 1
Library					kafka_2.12-2.4.0
Message					mac_address
Movies					macbook_list
Music					mei_card
MyJabberFiles				notes
Pictures				nxos.7.0.3.IHD8.0.703.bin
Postman					pack
Public					protobuf-3.15.1
Sites					python_test
Walkme					recovery_key
a					redis
aa					room
bazel-3.7.0-installer-darwin-x86_64.sh	stop_crash
bb					telemetry
bin					test
cc					tmp
contact					unit-test-3
LLUE-M-C5ME:~ llue$ cd goLang
LLUE-M-C5ME:goLang llue$ ls
hello
LLUE-M-C5ME:goLang llue$ cd hello/
LLUE-M-C5ME:hello llue$ ls
1		evtlog.proto	go.mod		go.sum		grpc-go		hello.go
LLUE-M-C5ME:hello llue$ cd grpc-go
LLUE-M-C5ME:grpc-go llue$ ls
AUTHORS					health
CODE-OF-CONDUCT.md			install_gae.sh
CONTRIBUTING.md				interceptor.go
Documentation				internal
GOVERNANCE.md				interop
LICENSE					keepalive
MAINTAINERS.md				metadata
Makefile				peer
README.md				picker_wrapper.go
SECURITY.md				picker_wrapper_test.go
attributes				pickfirst.go
backoff					pickfirst_test.go
backoff.go				preloader.go
balancer				profiling
balancer_conn_wrappers.go		reflection
balancer_conn_wrappers_test.go		regenerate.sh
balancer_switching_test.go		resolver
benchmark				resolver_conn_wrapper.go
binarylog				resolver_conn_wrapper_test.go
call.go					rpc_util.go
call_test.go				rpc_util_test.go
channelz				security
clientconn.go				server.go
clientconn_state_transition_test.go	server_test.go
clientconn_test.go			service_config.go
cmd					service_config_test.go
codec.go				serviceconfig
codec_test.go				stats
codegen.sh				status
codes					stream.go
connectivity				stress
credentials				tap
dialoptions.go				test
doc.go					testdata
encoding				trace.go
examples				trace_test.go
go.mod					version.go
go.sum					vet.sh
grpc_test.go				xds
grpclog
LLUE-M-C5ME:grpc-go llue$ cd ../
LLUE-M-C5ME:hello llue$ ls -ltr
total 40
-rw-r--r--   1 llue  staff    77 Mar 18 00:53 hello.go
-rw-r--r--   1 llue  staff  2280 Mar 18 01:00 go.sum
-rw-r--r--   1 llue  staff   161 Mar 18 01:00 go.mod
drwxr-xr-x  84 llue  staff  2688 Mar 18 01:01 grpc-go
-rw-r--r--   1 llue  staff   474 Mar 18 16:21 1
-rw-r--r--   1 llue  staff   474 Mar 18 16:21 evtlog.proto
LLUE-M-C5ME:hello llue$ cd ../
LLUE-M-C5ME:goLang llue$ ls
hello
LLUE-M-C5ME:goLang llue$ cd ../
LLUE-M-C5ME:~ llue$ ls
10.24.81.151				docker
2020-06-11.saz				evtlog
2020-07-07.saz				evtlog.pb.go
Applications				evtlog.proto
Calibre Library				fex_sensor.bin
Desktop					go
Documents				goLang
Downloads				grpc-go
Google Drive				iCloud Drive (Archive)
Header					iCloud Drive (Archive) - 1
Library					kafka_2.12-2.4.0
Message					mac_address
Movies					macbook_list
Music					mei_card
MyJabberFiles				notes
Pictures				nxos.7.0.3.IHD8.0.703.bin
Postman					pack
Public					protobuf-3.15.1
Sites					python_test
Walkme					recovery_key
a					redis
aa					room
bazel-3.7.0-installer-darwin-x86_64.sh	stop_crash
bb					telemetry
bin					test
cc					tmp
contact					unit-test-3
LLUE-M-C5ME:~ llue$ pwd
/Users/llue
LLUE-M-C5ME:~ llue$ vim evtlog.pb.go
LLUE-M-C5ME:~ llue$ ls -ltr evtlog.pb.go
-rw-r--r--  1 llue  staff  4739 Mar 16 18:19 evtlog.pb.go
LLUE-M-C5ME:~ llue$ pwd
/Users/llue
LLUE-M-C5ME:~ llue$ !protoc
protoc --go_out=. *.proto
protoc-gen-go: program not found or is not executable
Please specify a program using absolute path or make sure the program is available in your PATH system variable
--go_out: protoc-gen-go: Plugin failed with status code 1.
LLUE-M-C5ME:~ llue$ export GO111MODULE=on 
LLUE-M-C5ME:~ llue$ go get google.golang.org/protobuf/cmd/protoc-gen-go \
>          google.golang.org/grpc/cmd/protoc-gen-go-grpc
go: found google.golang.org/protobuf/cmd/protoc-gen-go in google.golang.org/protobuf v1.26.0
go: google.golang.org/grpc/cmd/protoc-gen-go-grpc upgrade => v1.1.0
LLUE-M-C5ME:~ llue$ export PATH="$PATH:$(go env GOPATH)/bin"
LLUE-M-C5ME:~ llue$ pwd
/Users/llue
LLUE-M-C5ME:~ llue$ ls -ltr
total 86696
drwxr-xr-x+    5 llue  staff       160 Oct 16  2015 Public
-rw-r--r--     1 llue  staff       633 Oct 29  2015 room
-rw-r--r--     1 llue  staff    421858 Nov  5  2015 bb
-rw-r--r--     1 llue  staff      1523 Nov  5  2015 cc
-rw-r--r--     1 llue  staff       882 Mar 25  2016 macbook_list
drwx------     4 llue  staff       128 May  7  2016 Applications
-rw-r--r--     1 llue  staff       200 May 19  2016 contact
-rw-r--r--     1 llue  staff         9 Sep  4  2016 recovery_key
drwxr-xr-x     3 llue  staff        96 Oct 13  2016 iCloud Drive (Archive)
-rw-r--r--     1 llue  staff         0 May  2  2018 fex_sensor.bin
-rw-r--r--     1 llue  staff         0 May  2  2018 nxos.7.0.3.IHD8.0.703.bin
drwxr-xr-x     2 llue  staff        64 Jun 27  2018 docker
-rw-r--r--     1 llue  staff        82 Sep 13  2018 mei_card
drwxr-xr-x+    3 llue  staff        96 Oct 13  2018 Sites
-rwxr-xr-x     1 llue  staff        72 Feb  5  2019 stop_crash
drwx------+    9 llue  staff       288 Mar 10  2019 Pictures
drwxr-xr-x     3 llue  staff        96 May 22  2019 Postman
-rw-r--r--     1 llue  staff      3836 Jul 27  2019 mac_address
drwxr-xr-x     5 llue  staff       160 Aug  9  2019 test
-rw-r--r--     1 llue  staff       340 Jan 12  2020 a
drwx------@   12 llue  staff       384 Mar  8  2020 Google Drive
drwxr-xr-x     4 llue  staff       128 Apr 12  2020 python_test
drwxr-xr-x    22 llue  staff       704 Apr 21  2020 Calibre Library
-rw-r--r--     1 llue  staff         0 Apr 24  2020 10.24.81.151
-rw-r--r--     1 llue  staff         0 May 24  2020 aa
-rw-r--r--     1 llue  staff    273105 Jun 11  2020 2020-06-11.saz
drwxr--r--     4 llue  staff       128 Jul  1  2020 MyJabberFiles
-rw-r--r--     1 llue  staff   2171699 Jul  7  2020 2020-07-07.saz
-rw-r--r--     1 llue  staff         0 Sep 22 23:24 pack
-rw-r--r--     1 llue  staff         0 Sep 22 23:24 Header
-rw-r--r--     1 llue  staff         0 Sep 22 23:24 Message
drwxr-xr-x     2 llue  staff        64 Oct 20 18:01 telemetry
-rw-r--r--     1 llue  staff      7835 Oct 22 03:43 unit-test-3
-rwxr-xr-x     1 llue  staff  40839901 Oct 23 12:41 bazel-3.7.0-installer-darwin-x86_64.sh
drwxr-xr-x     3 llue  staff        96 Oct 23 12:41 bin
-rw-r--r--     1 llue  staff      7279 Nov 24 20:18 notes
drwx------+    6 llue  staff       192 Dec 23 13:05 Music
drwx------     6 llue  staff       192 Dec 26 08:46 iCloud Drive (Archive) - 1
drwxr-xr-x@   10 llue  staff       320 Dec 26 08:46 kafka_2.12-2.4.0
drwxr-xr-x     5 llue  staff       160 Dec 26 08:46 redis
drwx------+    7 llue  staff       224 Dec 26 08:48 Movies
drwxr-xr-x     3 root  staff        96 Jan 14 09:40 Walkme
drwxr-xr-x@   45 llue  staff      1440 Feb 22 16:49 protobuf-3.15.1
-rw-r--r--     1 llue  staff      4739 Mar 16 18:19 evtlog.pb.go
drwxr-xr-x     3 llue  staff        96 Mar 18 00:51 goLang
drwxr-xr-x     4 llue  staff       128 Mar 18 01:00 go
drwxr-xr-x    84 llue  staff      2688 Mar 18 01:19 grpc-go
drwxr-xr-x     5 llue  staff       160 Mar 18 15:48 tmp
drwx------+   90 llue  staff      2880 Mar 19 11:19 Library
drwx------+ 1136 llue  staff     36352 Mar 19 21:39 Downloads
drwx------+  228 llue  staff      7296 Mar 20 02:36 Desktop
drwx------+  125 llue  staff      4000 Mar 20 17:21 Documents
drwxr-xr-x     2 llue  staff        64 Mar 21 14:06 evtlog
-rw-r--r--     1 llue  staff       414 Mar 21 14:08 evtlog.proto
LLUE-M-C5ME:~ llue$ protoc --go_out=. *.proto
protoc-gen-go: invalid Go import path "evtlog" for "evtlog.proto"

The import path must contain at least one forward slash ('/') character.

See https://developers.google.com/protocol-buffers/docs/reference/go-generated#package for more information.

--go_out: protoc-gen-go: Plugin failed with status code 1.
LLUE-M-C5ME:~ llue$ vim evtlog.proto 
LLUE-M-C5ME:~ llue$ protoc --go_out=. *.proto
LLUE-M-C5ME:~ llue$ ls -ltr
total 86696
drwxr-xr-x+    5 llue  staff       160 Oct 16  2015 Public
-rw-r--r--     1 llue  staff       633 Oct 29  2015 room
-rw-r--r--     1 llue  staff    421858 Nov  5  2015 bb
-rw-r--r--     1 llue  staff      1523 Nov  5  2015 cc
-rw-r--r--     1 llue  staff       882 Mar 25  2016 macbook_list
drwx------     4 llue  staff       128 May  7  2016 Applications
-rw-r--r--     1 llue  staff       200 May 19  2016 contact
-rw-r--r--     1 llue  staff         9 Sep  4  2016 recovery_key
drwxr-xr-x     3 llue  staff        96 Oct 13  2016 iCloud Drive (Archive)
-rw-r--r--     1 llue  staff         0 May  2  2018 fex_sensor.bin
-rw-r--r--     1 llue  staff         0 May  2  2018 nxos.7.0.3.IHD8.0.703.bin
drwxr-xr-x     2 llue  staff        64 Jun 27  2018 docker
-rw-r--r--     1 llue  staff        82 Sep 13  2018 mei_card
drwxr-xr-x+    3 llue  staff        96 Oct 13  2018 Sites
-rwxr-xr-x     1 llue  staff        72 Feb  5  2019 stop_crash
drwx------+    9 llue  staff       288 Mar 10  2019 Pictures
drwxr-xr-x     3 llue  staff        96 May 22  2019 Postman
-rw-r--r--     1 llue  staff      3836 Jul 27  2019 mac_address
drwxr-xr-x     5 llue  staff       160 Aug  9  2019 test
-rw-r--r--     1 llue  staff       340 Jan 12  2020 a
drwx------@   12 llue  staff       384 Mar  8  2020 Google Drive
drwxr-xr-x     4 llue  staff       128 Apr 12  2020 python_test
drwxr-xr-x    22 llue  staff       704 Apr 21  2020 Calibre Library
-rw-r--r--     1 llue  staff         0 Apr 24  2020 10.24.81.151
-rw-r--r--     1 llue  staff         0 May 24  2020 aa
-rw-r--r--     1 llue  staff    273105 Jun 11  2020 2020-06-11.saz
drwxr--r--     4 llue  staff       128 Jul  1  2020 MyJabberFiles
-rw-r--r--     1 llue  staff   2171699 Jul  7  2020 2020-07-07.saz
-rw-r--r--     1 llue  staff         0 Sep 22 23:24 pack
-rw-r--r--     1 llue  staff         0 Sep 22 23:24 Header
-rw-r--r--     1 llue  staff         0 Sep 22 23:24 Message
drwxr-xr-x     2 llue  staff        64 Oct 20 18:01 telemetry
-rw-r--r--     1 llue  staff      7835 Oct 22 03:43 unit-test-3
-rwxr-xr-x     1 llue  staff  40839901 Oct 23 12:41 bazel-3.7.0-installer-darwin-x86_64.sh
drwxr-xr-x     3 llue  staff        96 Oct 23 12:41 bin
-rw-r--r--     1 llue  staff      7279 Nov 24 20:18 notes
drwx------+    6 llue  staff       192 Dec 23 13:05 Music
drwx------     6 llue  staff       192 Dec 26 08:46 iCloud Drive (Archive) - 1
drwxr-xr-x@   10 llue  staff       320 Dec 26 08:46 kafka_2.12-2.4.0
drwxr-xr-x     5 llue  staff       160 Dec 26 08:46 redis
drwx------+    7 llue  staff       224 Dec 26 08:48 Movies
drwxr-xr-x     3 root  staff        96 Jan 14 09:40 Walkme
drwxr-xr-x@   45 llue  staff      1440 Feb 22 16:49 protobuf-3.15.1
-rw-r--r--     1 llue  staff      4739 Mar 16 18:19 evtlog.pb.go
drwxr-xr-x     3 llue  staff        96 Mar 18 00:51 goLang
drwxr-xr-x     4 llue  staff       128 Mar 18 01:00 go
drwxr-xr-x    84 llue  staff      2688 Mar 18 01:19 grpc-go
drwxr-xr-x     5 llue  staff       160 Mar 18 15:48 tmp
drwx------+   90 llue  staff      2880 Mar 19 11:19 Library
drwx------+ 1136 llue  staff     36352 Mar 19 21:39 Downloads
drwx------+  228 llue  staff      7296 Mar 20 02:36 Desktop
drwx------+  125 llue  staff      4000 Mar 20 17:21 Documents
-rw-r--r--     1 llue  staff       416 Mar 21 14:20 evtlog.proto
drwxr-xr-x     3 llue  staff        96 Mar 21 14:20 evtlog
LLUE-M-C5ME:~ llue$ cd evtlog
LLUE-M-C5ME:evtlog llue$ ls
evtlog.pb.go
LLUE-M-C5ME:evtlog llue$ ls -ltr
total 16
-rw-r--r--  1 llue  staff  6739 Mar 21 14:20 evtlog.pb.go
LLUE-M-C5ME:evtlog llue$ vim evtlog.pb.go

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
//      protoc-gen-go v1.26.0
//      protoc        v3.12.3
// source: evtlog.proto

package evtlog

import (
        protoreflect "google.golang.org/protobuf/reflect/protoreflect"
        protoimpl "google.golang.org/protobuf/runtime/protoimpl"
        reflect "reflect"
        sync "sync"
)

const (
        // Verify that this generated code is sufficiently up-to-date.
        _ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
        // Verify that runtime/protoimpl is sufficiently up-to-date.
        _ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Evtlog struct {
        state         protoimpl.MessageState
        sizeCache     protoimpl.SizeCache
        unknownFields protoimpl.UnknownFields

        Version    string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
        FileType   string `protobuf:"bytes,2,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`
        Compressed bool   `protobuf:"varint,3,opt,name=compressed,proto3" json:"compressed,omitempty"`
        SeqNo      uint64 `protobuf:"varint,4,opt,name=seq_no,json=seqNo,proto3" json:"seq_no,omitempty"`
        FileName   string `protobuf:"bytes,5,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
        Size       uint64 `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
        ChunkTotal uint32 `protobuf:"varint,7,opt,name=chunk_total,json=chunkTotal,proto3" json:"chunk_total,omitempty"`
        ChunkNo    uint32 `protobuf:"varint,8,opt,name=chunk_no,json=chunkNo,proto3" json:"chunk_no,omitempty"`
        File       []byte `protobuf:"bytes,9,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *Evtlog) Reset() {
        *x = Evtlog{}
        if protoimpl.UnsafeEnabled {
                mi := &file_evtlog_proto_msgTypes[0]
                ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
                ms.StoreMessageInfo(mi)
        }
}

func (x *Evtlog) String() string {

