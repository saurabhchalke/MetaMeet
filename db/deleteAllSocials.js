import faunadb from "faunadb";

const q = faunadb.query;

const client = new faunadb.Client({
  secret: process.env.FAUNADB_SERVER_SECRET,
});

// Delete all socials.
const response = await client.query(
  q.Map(
    q.Paginate(q.Documents(q.Collection("social"))),
    q.Lambda("X", q.Delete(q.Var("X")))
  )
);
console.info(`Response: ${JSON.stringify(response, null, 2)}`);
