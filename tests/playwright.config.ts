import { defineConfig } from "@playwright/test";

export default defineConfig({
  testDir: ".",
  testMatch: "integration.spec.ts",
  use: {
    baseURL: process.env.PLAYWRIGHT_TEST_BASE_URL || "http://localhost:3000",
    trace: "on-first-retry",
    headless: process.env.CI ? true : false,
    launchOptions: {
      slowMo: 1000,
    },
  },
  timeout: 30000,
  reporter: "list",
});
