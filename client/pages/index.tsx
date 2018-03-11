import * as React from 'react'
import Head from 'next/head'
import fetch from 'isomorphic-fetch'

import getConfig from 'next/config'

export default class GraphiQLPage extends React.Component<{ graphqlEndpoint: string }, {}> {
  static getInitialProps(ctx) {
    const { publicRuntimeConfig } = getConfig()
    return {
      graphqlEndpoint: publicRuntimeConfig.graphqlEndpoint
    }
  }

  graphQLFetcher = endpoint => graphQLParams => {
    return fetch(endpoint, {
      method: 'post',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(graphQLParams)
    }).then(response => response.json())
  }
  render() {
    return (
      <div>
        {this.props.graphqlEndpoint}
      </div>
    )
  }
}
