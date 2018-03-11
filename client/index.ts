require('dotenv').config({path: '../.env'})
import * as next from 'next'
import * as express from 'express'

async function initServer() {
  const server = express()
  const isProd = process.env.NODE_ENV === 'production'

  const client = next({ dev: !isProd })
  await client.prepare()

  if(!isProd) {
    const { graphiqlExpress } = require('apollo-server-express');
    const endpointURL = process.env.GRAPHQL_ENDPOINT || '/graphql'
    console.log(`[server] init graphiql@${endpointURL}`)
    server.get('/graphiql', graphiqlExpress({ endpointURL  })); // if you want GraphiQL enabled
  }
  server.use(client.getRequestHandler())
  server.listen(3000, function() {
    console.log('[index] start client and source at ' + 3000)
  })
}

initServer()
