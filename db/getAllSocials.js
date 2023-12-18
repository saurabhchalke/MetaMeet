import faunadb from "faunadb";

const q = faunadb.query;

const client = new faunadb.Client({
  domain: process.env.FAUNADB_DOMAIN,
  port: process.env.FAUNADB_PORT,
  scheme: process.env.FAUNADB_SCHEME,
  secret: process.env.FAUNADB_SERVER_SECRET,
});

const response = await client.query(q.Map(q.Paginate(q.Documents(q.Collection("social"))), q.Lambda("X", q.Get(q.Var("X")))));
console.info(`Response: ${JSON.stringify(response, null, 2)}`);
