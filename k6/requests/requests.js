import { statusCheck } from "../asserts/checks.js";
import http from "k6/http";
import { check } from "k6";

// Function to make API requests
export function makeRequests(appURL, endpointList) {
  const responses = {};

  for (let endpoint in endpointList) {
    if (endpointList.hasOwnProperty(endpoint)) {
      const apiUrl = `${appURL}${endpointList[endpoint]}`;
      const response = http.get(apiUrl);
      responses[endpoint] = response;
    }
  }

  return responses;
}

// Function to check responses
export function checkResponses(responses) {
  for (let endpoint in responses) {
    if (responses.hasOwnProperty(endpoint)) {
      const response = responses[endpoint];
      check(response, statusCheck);
    }
  }
}
