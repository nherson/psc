import { createPromiseClient } from "@bufbuild/connect";
import { createConnectTransport } from "@bufbuild/connect-web";
import { PSCService } from "./psc_connect";
const baseUrl = "https://psc-fantasy.fly.dev/api";

const transport = createConnectTransport({
  baseUrl: baseUrl,
});

export const client = createPromiseClient(PSCService, transport);
