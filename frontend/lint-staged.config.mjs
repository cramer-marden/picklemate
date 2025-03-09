/**
 * @filename: lint-staged.config.js
 * @type {import('lint-staged').Configuration}
 */
export default {
  // Run ESLint on JS, TS, and TSX files
  '**/*.{js,mjs,cjs,jsx,ts,tsx}': [
    // Only lint files that are staged (not the entire codebase)
    'bun run lint:fix',
    // Format the code after linting
    'bun run format',
  ],
  // Run Prettier on other file types
  '**/*.{json,md}': ['prettier --write'],
};
