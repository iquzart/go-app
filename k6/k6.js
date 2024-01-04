import http from "k6/http";
import { sleep } from "k6";
import { check } from "k6";
import { endpointList } from "./endpoints/endpoints.js";
import { makeRequests, checkResponses } from "./requests/requests.js";

export const options = {
  discardResponseBodies: true,
  scenarios: {
    contacts: {
      executor: "ramping-vus",
      startVUs: 0,
      stages: [
        { target: 50, duration: "30s" }, // linearly go from 3 VUs to 50 VUs for 30s
        { target: 200, duration: "10s" }, // linearly go from 50 VUs to 200 VUs for 10s
        { target: 200, duration: "5m" }, // continue with 200 VUs for 5 minutes
      ],
      gracefulRampDown: "0s",
    },
  },
  thresholds: {
    http_req_failed: ["rate<0.01"], // http errors should be less than 1%
    http_req_duration: ["p(95)<200"], // 95% of requests should be below 200ms
  },
};

export default function () {
  const appDomain = "go-app.testbox.pod";

  const urlHTTPScheme = "https";

  const appURL = `${urlHTTPScheme}://${appDomain}`;

  const responses = makeRequests(appURL, endpointList);
  checkResponses(responses);

  sleep(1);
}
