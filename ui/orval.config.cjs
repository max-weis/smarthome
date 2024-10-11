module.exports = {
  'device-api': {
    input: '../device/api.yaml',
    output: {
      mode: 'split',
      target: 'src/api/device-api.ts',
      client: 'axios',
      override: {
        mutator: {
          path: './src/api/axios-instance.ts',
          name: 'customInstance',
        },
      },
    },
    hooks: {
      afterAllFilesWrite: 'prettier --write',
    },
  },
};