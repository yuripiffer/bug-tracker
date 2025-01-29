import { test, expect } from "@playwright/test";

test("Call Health Check", async ({ request }) => {
  console.log("Starting test...");

  const healthCheckResponse = await request.get("health");
  expect(healthCheckResponse.ok()).toBeTruthy();
  expect(await healthCheckResponse.json()).toEqual(
    expect.objectContaining({
      status: "ok",
    })
  );
});
