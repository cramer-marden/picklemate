/**
 * @filename: lint-staged.config.js
 * @type {import('lint-staged').Configuration}
 */
export default {
  // Run ESLint on JS, TS, and TSX files
  "**/*.{js,jsx,ts,tsx}": [
    // Only lint files that are staged (not the entire codebase)
    'eslint --fix',
    // Format the code after linting
    'prettier --write',
  ],
  // Run Prettier on other file types
  '**/*.{json,md}': ["prettier --write"],
};
