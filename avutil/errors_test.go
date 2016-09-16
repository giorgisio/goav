package avutil

import (
	"fmt"
	"testing"
)

func TestPrintErrors(t *testing.T) {
	fmt.Printf("%s: %d | %s\n", "AVERROR_BSF_NOT_FOUND", AVERROR_BSF_NOT_FOUND(), AvStrerror(AVERROR_BSF_NOT_FOUND()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_BUG", AVERROR_BUG(), AvStrerror(AVERROR_BUG()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_BUFFER_TOO_SMALL", AVERROR_BUFFER_TOO_SMALL(), AvStrerror(AVERROR_BUFFER_TOO_SMALL()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_DECODER_NOT_FOUND", AVERROR_DECODER_NOT_FOUND(), AvStrerror(AVERROR_DECODER_NOT_FOUND()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_DEMUXER_NOT_FOUND", AVERROR_DEMUXER_NOT_FOUND(), AvStrerror(AVERROR_DEMUXER_NOT_FOUND()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_ENCODER_NOT_FOUND", AVERROR_ENCODER_NOT_FOUND(), AvStrerror(AVERROR_ENCODER_NOT_FOUND()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_EOF", AVERROR_EOF(), AvStrerror(AVERROR_EOF()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_EXIT", AVERROR_EXIT(), AvStrerror(AVERROR_EXIT()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_EXTERNAL", AVERROR_EXTERNAL(), AvStrerror(AVERROR_EXTERNAL()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_FILTER_NOT_FOUND", AVERROR_FILTER_NOT_FOUND(), AvStrerror(AVERROR_FILTER_NOT_FOUND()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_INVALIDDATA", AVERROR_INVALIDDATA(), AvStrerror(AVERROR_INVALIDDATA()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_MUXER_NOT_FOUND", AVERROR_MUXER_NOT_FOUND(), AvStrerror(AVERROR_MUXER_NOT_FOUND()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_OPTION_NOT_FOUND", AVERROR_OPTION_NOT_FOUND(), AvStrerror(AVERROR_OPTION_NOT_FOUND()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_PATCHWELCOME", AVERROR_PATCHWELCOME(), AvStrerror(AVERROR_PATCHWELCOME()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_PROTOCOL_NOT_FOUND", AVERROR_PROTOCOL_NOT_FOUND(), AvStrerror(AVERROR_PROTOCOL_NOT_FOUND()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_STREAM_NOT_FOUND", AVERROR_STREAM_NOT_FOUND(), AvStrerror(AVERROR_STREAM_NOT_FOUND()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_UNKNOWN", AVERROR_UNKNOWN(), AvStrerror(AVERROR_UNKNOWN()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_EXPERIMENTAL", AVERROR_EXPERIMENTAL(), AvStrerror(AVERROR_EXPERIMENTAL()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_INPUT_CHANGED", AVERROR_INPUT_CHANGED(), AvStrerror(AVERROR_INPUT_CHANGED()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_OUTPUT_CHANGED", AVERROR_OUTPUT_CHANGED(), AvStrerror(AVERROR_OUTPUT_CHANGED()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_HTTP_BAD_REQUEST", AVERROR_HTTP_BAD_REQUEST(), AvStrerror(AVERROR_HTTP_BAD_REQUEST()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_HTTP_UNAUTHORIZED", AVERROR_HTTP_UNAUTHORIZED(), AvStrerror(AVERROR_HTTP_UNAUTHORIZED()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_HTTP_FORBIDDEN", AVERROR_HTTP_FORBIDDEN(), AvStrerror(AVERROR_HTTP_FORBIDDEN()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_HTTP_NOT_FOUND", AVERROR_HTTP_NOT_FOUND(), AvStrerror(AVERROR_HTTP_NOT_FOUND()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_HTTP_OTHER_4XX", AVERROR_HTTP_OTHER_4XX(), AvStrerror(AVERROR_HTTP_OTHER_4XX()))
	fmt.Printf("%s: %d | %s\n", "AVERROR_HTTP_SERVER_ERROR", AVERROR_HTTP_SERVER_ERROR(), AvStrerror(AVERROR_HTTP_SERVER_ERROR()))
	fmt.Printf("%s: %d\n", "AV_ERROR_MAX_STRING_SIZE", AV_ERROR_MAX_STRING_SIZE())
}
