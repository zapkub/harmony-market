require('dotenv').config({path: '../.env'})
const isProd = process.env.NODE_ENV === 'production'

module.exports = {
  publicRuntimeConfig: { // Will be available on both server and client
    graphqlEndpoint: process.env.GRAPHQL_ENDPOINT || '/graphql',
    dev: !isProd
  }
}