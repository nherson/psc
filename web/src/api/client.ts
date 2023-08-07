import { createPromiseClient } from "@bufbuild/connect";
import { createConnectTransport } from "@bufbuild/connect-web";
import { PSCService } from "./psc_connect";
export const baseUrl = "https://poopswagchampionship.com/api";

const transport = createConnectTransport({
  baseUrl: baseUrl,
});

export const client = createPromiseClient(PSCService, transport);
