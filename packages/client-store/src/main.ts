/**
 * main.ts
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Components
import App from './App.vue'

// Composables
import { createApp } from 'vue'

// Plugins
import { registerPlugins } from '@/plugins'

const app = createApp(App)

registerPlugins(app)

import {
  createConnectTransport,
  createPromiseClient
} from "@bufbuild/connect-web";
import { GreetService } from "@ticketbooth/server/proto/store/v1/store_connectweb";

const transport = createConnectTransport({
  baseUrl: "http://localhost:9100",
});

const client = createPromiseClient(GreetService, transport);

async function doer() {
  const res = await client.greet({name: "foo"})
  console.log(res)
}
doer()

app.mount('#app')
