package node

import (
	"bytes"
	"fmt"
	"github.com/donkeysharp/gocho/pkg/config"
	"net/http"
	"regexp"
)

const (
	HTML_BODY = `<html>
<head>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css" type="text/css">
    <style>
        * {
            font-family: sans-serif;
        }
        a {
            text-decoration: none;
            color: #1552A8;
            display: block;
            padding-bottom: 3px;
        }
        a.directory:before {
            color: #FFCF30;
            font-family: FontAwesome;
            content: "\f07b\00a0";
        }
        a.file:before {
            font-family: FontAwesome;
            content: "\f016\00a0";
        }
    </style>
    <script>
        function goBack() {
            var path = window.location.pathname.split('/');
            if (path.length <= 3) {
                window.location = '/';
                return;
            }
            window.location = path.slice(0, path.length - 2).join('/');
        }
    </script>
</head>
<body>
<a class="directory" onClick="javascript:goBack()" href="#">..</a>`
	HTML_END = `<script type="text/javascript">
    document.addEventListener('DOMContentLoaded', () => {
        document.querySelectorAll(".file").forEach((el) => {
            document.body.appendChild(el);
        })
    })
</script>
</body>
</html>`
)

type FileServerResponseInterceptor struct {
	OriginalWriter http.ResponseWriter
	IndexBuffer    *bytes.Buffer
}

func (f *FileServerResponseInterceptor) WriteHeader(status int) {
	f.OriginalWriter.WriteHeader(status)
}

func (f *FileServerResponseInterceptor) Header() http.Header {
	return f.OriginalWriter.Header()
}

func (f *FileServerResponseInterceptor) Write(content []byte) (int, error) {
	// if it's not an html tag why bother evaluating with regex?
	if content[0] != byte('<') {
		return f.OriginalWriter.Write(content)
	}
	re := regexp.MustCompile("^<a.+href=\"(.+)\".*>(.+)</a>$|^</{0,1}pre>$")
	if !re.Match(bytes.Trim(content, "\n\r")) {
		return f.OriginalWriter.Write(content)
	}
	content = bytes.Trim(content, "\n\r")
	directoryRegex := regexp.MustCompile("^<a.+href=\"(.+/)\".*>(.+)</a>$")
	if directoryRegex.Match(content) {
		directoryLink := "<a class=\"directory\" href=\"$1\">$2</a>\n"
		content = directoryRegex.ReplaceAll(content, []byte(directoryLink))
		return f.IndexBuffer.Write(content)
	}
	fileRegex := regexp.MustCompile("^<a.+href=\"(.+)\".*>(.+)</a>$")
	if fileRegex.Match(content) {
		fileLink := "<a class=\"file\" href=\"$1\">$2</a>\n"
		content = fileRegex.ReplaceAll(content, []byte(fileLink))
		return f.IndexBuffer.Write(content)
	}
	return 0, nil
}

func interceptorHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		interceptor := &FileServerResponseInterceptor{
			OriginalWriter: w,
			IndexBuffer:    bytes.NewBuffer(nil),
		}
		// r.Header.Del("If-Modified-Since")
		next.ServeHTTP(interceptor, r)

		if interceptor.IndexBuffer.Len() > 0 {
			w.Write([]byte(HTML_BODY))
			w.Write(interceptor.IndexBuffer.Bytes())
			w.Write([]byte(HTML_END))
		}
	}
	return http.HandlerFunc(fn)
}

func fileServe(conf *config.Config) {
	fileMux := http.NewServeMux()
	fileMux.Handle("/", interceptorHandler(http.FileServer(http.Dir(conf.ShareDirectory))))
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", conf.WebPort), fileMux)
}
