module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: [
    'plugin:vue/essential',
    'eslint:recommended',
    '@vue/typescript/recommended',
    '@vue/prettier',
    '@vue/prettier/@typescript-eslint',
  ],
  parserOptions: {
    ecmaVersion: 2020,
  },
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'vue/no-use-v-if-with-v-for': ['off'],
    'vue/experimental-script-setup-vars': ['off'],
    'vue/component-name-in-template-casing': ['off'],
    'vue/no-unused-components': ['off'],
    'prettier/prettier': [
      'error',
      {
        htmlWhitespaceSensitivity: 'ignore',
        semi: true,
        singleQuote: true,
      },
    ],
    'vue/html-self-closing': [
      'error',
      {
        html: {
          void: 'any',
        },
      },
    ],
  },
  overrides: [
    {
      files: [
        '**/__tests__/*.{j,t}s?(x)',
        '**/tests/unit/**/*.spec.{j,t}s?(x)',
      ],
      env: {
        jest: true,
      },
    },
  ],
};
