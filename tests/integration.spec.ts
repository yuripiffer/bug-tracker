import { test, expect } from "@playwright/test";

test("Bug creation and deletion flow", async ({ page }) => {
  // Add console log listener
  page.on("console", (msg) => {
    console.log(`Browser console [${msg.type()}]: ${msg.text()}`);
  });

  console.log("Starting test...");

  // Visit the homepage
  await page.goto("/");
  console.log("Navigated to homepage");

  // Debug: Log the page content
  console.log("Page content:", await page.content());

  // Wait for page to be ready
  await page.waitForLoadState("networkidle");
  console.log("Page loaded");

  // Wait for the button with more specific selector
  await page.waitForSelector('button:text("Add New Bug")', { timeout: 60000 });
  console.log("Found Add New Bug button");

  // Click on the "Add Bug" button
  await page.click('button:text("Add New Bug")');
  console.log("Clicked Add New Bug button");

  // Wait for the modal to be visible
  console.log("Waiting for modal...");
  await page.waitForSelector('[data-testid="add-bug-modal"]', {
    timeout: 60000,
  });
  console.log("Modal found");

  await page.waitForSelector('[data-testid="add-bug-form"]', {
    timeout: 60000,
  });
  console.log("Form found");

  // Debug: Log modal content
  const modalContent = await page
    .locator('[data-testid="add-bug-modal"]')
    .innerHTML();
  console.log("Modal content:", modalContent);

  // Fill out the form and submit
  await page.fill('input[name="title"]', "Test Bug");
  await page.fill('textarea[name="description"]', "This is a test bug");
  await page.selectOption('select[name="priority"]', "Medium");
  await page.click('button:text("Add Bug")');

  // Verify that the bug was created
  await expect(page.locator('a:text("Test Bug")').first()).toBeVisible();

  // Click on the bug to view details
  await page.locator('a:text("Test Bug")').first().click();

  // Verify the bug details
  await expect(page.locator("h1:text('Test Bug')")).toBeVisible();
  await expect(page.locator("text=This is a test bug")).toBeVisible();
  await expect(page.locator("text=Medium")).toBeVisible();

  // Delete the bug
  await page.click("text=Delete Bug");
  await page.click("text=Confirm");

  // Verify the bug was deleted
  await expect(page.locator('a:text("Test Bug")').first()).toBeHidden();
});

test("Adding a comment to a bug", async ({ page }) => {
  console.log("Starting comment test...");

  // Visit the homepage
  await page.goto("/");
  console.log("Navigated to homepage");

  // Create a bug first
  await page.click('button:text("Add New Bug")');
  await page.waitForSelector('[data-testid="add-bug-form"]', {
    timeout: 60000,
  });
  await page.fill('input[name="title"]', "Bug for Comment");
  await page.fill('textarea[name="description"]', "This bug needs a comment");
  await page.selectOption('select[name="priority"]', "Medium");
  await page.click('button:text("Add Bug")');

  // Wait for bug to be created and click on it
  await expect(page.locator('a:text("Bug for Comment")').first()).toBeVisible();
  await page.locator('a:text("Bug for Comment")').first().click();

  // Wait for comment form
  await page.waitForSelector('[data-testid="comment-form"]', {
    timeout: 60000,
  });
  console.log("Comment form found");

  // Add a comment
  const timestamp = "test-comment-" + Math.random().toString(36).substring(7);
  await page.fill(
    '[data-testid="comment-content"]',
    `Test comment ${timestamp}`
  );
  await page.fill('[data-testid="comment-author"]', `Test User ${timestamp}`);
  await page.click('button:text("Add Comment")');

  // Verify the comment was added - with longer timeout
  await expect(page.locator(`p:text('Test comment ${timestamp}')`)).toBeVisible(
    { timeout: 10000 }
  ); // Increase timeout and use static string
});

test("Editing a bug", async ({ page }) => {
  console.log("Starting edit bug test...");

  // Create a bug
  await page.goto("/");
  await page.click('button:text("Add New Bug")');
  await page.waitForSelector('[data-testid="add-bug-form"]', {
    timeout: 60000,
  });
  await page.fill('input[name="title"]', "Bug to Edit");
  await page.fill('textarea[name="description"]', "This bug will be edited");
  await page.selectOption('select[name="priority"]', "Low");
  await page.click('button:text("Add Bug")');

  // Navigate to bug details
  await page.locator('a:text("Bug to Edit")').first().click();

  // Click on edit button
  await page.click('button:text("Edit Bug")');
  console.log("Clicked edit button");

  // Wait for edit form
  await page.waitForSelector('[data-testid="edit-bug-form"]', {
    timeout: 60000,
  });
  console.log("Edit form found");

  // Debug: Log form content
  const formContent = await page
    .locator('[data-testid="edit-bug-form"]')
    .innerHTML();
  console.log("Edit form content:", formContent);

  // Update bug details
  await page.fill('input[name="title"]', "Edited Bug Title");
  await page.fill('textarea[name="description"]', "This bug has been edited");
  await page.selectOption('select[name="priority"]', "High");
  await page.click('button:text("Save Changes")');

  // Verify updated details
  await expect(page.locator("h1:text('Edited Bug Title')")).toBeVisible();
  await expect(page.locator("text=This bug has been edited")).toBeVisible();
  await expect(page.locator("text=High")).toBeVisible();
});
