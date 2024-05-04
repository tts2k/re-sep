import protoLoader from "@grpc/proto-loader";
import { loadPackageDefinition } from "@grpc/grpc-js";

export const packageDefinition = protoLoader.loadSync(
	"./src/lib/proto/content.proto",
	{
		keepCase: true,
		longs: String,
		enums: String,
		defaults: true,
		oneofs: true,
	},
);

export const proto = loadPackageDefinition(packageDefinition);
