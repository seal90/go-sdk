package virtual

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req,
		reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption) error {

		ctx = metadataProxy(ctx)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// StreamClientInterceptor
func StreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		callOpts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		ctx = metadataProxy(ctx)

		return streamer(ctx, desc, cc, method)
	}
}

func metadataProxy(ctx context.Context) context.Context {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if ns, exists := md["virtual-namespace"]; exists {
			ctx = metadata.AppendToOutgoingContext(ctx, "virtual-namespace", ns[0])
		}

	}
	return ctx
}
