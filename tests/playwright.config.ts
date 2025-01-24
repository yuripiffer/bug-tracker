import { defineConfig } from "@playwright/test";

export default defineConfig({
  testDir: ".",
  testMatch: "**/*.spec.ts",
  use: {
    baseURL: "http://localhost:3000",
    trace: "on-first-retry",
    headless: process.env.CI ? true : false,
    launchOptions: {
      slowMo: 1000,
    },
  },
});
