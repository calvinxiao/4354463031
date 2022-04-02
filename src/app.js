const axios = require("axios").default;
const Koa = require("koa");
const app = new Koa();

const { flag1 } = process.env;

const client = axios.create({
  baseURL: "http://127.0.0.1:8080/users/",
});

app.use(async (ctx) => {
  const { user_id } = ctx.query;
  console.log(`user_id is ${user_id}`)
  if (!user_id || Array.isArray(user_id)) {
    ctx.status = 404;
    return;
  }

  const url = `${user_id}`;
  try {
    const response = await client({
      headers: {
        "api-key": flag1,
      },
      method: "get",
      url,
    });
    ctx.status = 200;
    ctx.body = response.data;
  } catch (err) {
    console.error(err);
    throw err;
  }
});

app.listen(9003);
