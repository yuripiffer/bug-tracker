import { defineConfig } from "@playwright/test";

export default defineConfig({
  testDir: ".",
  testMatch: "integration.spec.ts",
  use: {
    baseURL: process.env.PLAYWRIGHT_TEST_BASE_URL || "http://localhost:3000",
    trace: "on-first-retry",
    headless: process.env.CI ? true : false,
    launchOptions: {
      slowMo: process.env.CI ? 0 : 1000,
    },
  },
  timeout: 30000,
  reporter: [
    ["list"],
    ["junit", { outputFile: "test-results.xml" }],
    ["html", { outputFolder: "playwright-report" }],
  ],
  webServer: process.env.CI
    ? undefined
    : {
        command: "npm run dev",
        url: "http://localhost:3000",
        reuseExistingServer: !process.env.CI,
      },
});
