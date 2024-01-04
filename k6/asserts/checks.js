// Note: You can customize or expand this object to include additional checks based on your specific testing requirements.

// This statusCheck object defines a set of criteria to evaluate the response received during load testing.

export const statusCheck = {
  "is status 200": (r) => r.status === 200, // Checks if the HTTP status code is 200 (OK)
  "protocol is HTTP/2": (r) => r.proto === "HTTP/2.0", // Verifies if the protocol used is HTTP/2
};
