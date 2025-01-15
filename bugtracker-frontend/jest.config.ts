import nextJest from 'next/jest'

const createJestConfig = nextJest({
    dir: './',
})

const config = {
    setupFilesAfterEnv: ['<rootDir>/jest.setup.js'],
    testEnvironment: 'jest-environment-jsdom',
    moduleNameMapper: {
        '^@/components/(.*)$': '<rootDir>/components/$1',
        '^@/pages/(.*)$': '<rootDir>/pages/$1',
    },
}

export default createJestConfig(config) 