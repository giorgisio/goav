#ifndef ERROR_MACROS_H
#define ERROR_MACROS_H

int macro_AVERROR(int e);

extern const int macro_AVERROR_BSF_NOT_FOUND;
extern const int macro_AVERROR_BUG;
extern const int macro_AVERROR_BUFFER_TOO_SMALL;
extern const int macro_AVERROR_DECODER_NOT_FOUND;
extern const int macro_AVERROR_DEMUXER_NOT_FOUND;
extern const int macro_AVERROR_ENCODER_NOT_FOUND;
extern const int macro_AVERROR_EOF;
extern const int macro_AVERROR_EXIT;
extern const int macro_AVERROR_EXTERNAL;
extern const int macro_AVERROR_FILTER_NOT_FOUND;
extern const int macro_AVERROR_INVALIDDATA;
extern const int macro_AVERROR_MUXER_NOT_FOUND;
extern const int macro_AVERROR_OPTION_NOT_FOUND;
extern const int macro_AVERROR_PATCHWELCOME;
extern const int macro_AVERROR_PROTOCOL_NOT_FOUND;
extern const int macro_AVERROR_STREAM_NOT_FOUND;
extern const int macro_AVERROR_UNKNOWN;
extern const int macro_AVERROR_EXPERIMENTAL;
extern const int macro_AVERROR_INPUT_CHANGED;
extern const int macro_AVERROR_OUTPUT_CHANGED;
extern const int macro_AVERROR_HTTP_BAD_REQUEST;
extern const int macro_AVERROR_HTTP_UNAUTHORIZED;
extern const int macro_AVERROR_HTTP_FORBIDDEN;
extern const int macro_AVERROR_HTTP_NOT_FOUND;
extern const int macro_AVERROR_HTTP_OTHER_4XX;
extern const int macro_AVERROR_HTTP_SERVER_ERROR;
extern const int macro_AV_ERROR_MAX_STRING_SIZE;

#endif
