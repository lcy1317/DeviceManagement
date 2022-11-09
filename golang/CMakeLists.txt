cmake_minimum_required(VERSION 3.7)

if (NOT GOPATH)
  set(GOPATH "${CMAKE_CURRENT_BINARY_DIR}/go")
endif()

add_custom_target(swagger-gen)
function(run_swagger_gen sub_project)
  add_custom_command (
    WORKING_DIRECTORY "${CMAKE_CURRENT_SOURCE_DIR}/${sub_project}"
    TARGET swagger-gen
    COMMAND ${SWAGGER_EXECUTABLE} -q generate server --exclude-main
  )
endfunction()
run_swagger_gen(lcyutils)

add_custom_target(go-cmd ALL DEPENDS swagger-gen)
add_custom_command(
  TARGET go-cmd
  WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}
  COMMAND ${CMAKE_COMMAND} -E make_directory ${CMAKE_CURRENT_BINARY_DIR}/bin
  COMMAND ${CMAKE_COMMAND} -E env
    "GOPATH=${GOPATH}"
    "CGO_ENABLED=0"
    ${GO_EXECUTABLE} build -installsuffix cgo -ldflags '-s' -o ${CMAKE_CURRENT_BINARY_DIR}/bin/ ./cmd/...
)

add_test(
  NAME go-test
  WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}
  COMMAND ${CMAKE_COMMAND} -E env
    ${GO_EXECUTABLE} test ./...
)

set_property(TEST go-test APPEND PROPERTY DEPENDS swagger-gen)
