const axios = require("axios");

// Fetch all data of the current type that points to the previous type with the given id
export default async function fetchData({ curr, prev, prevId }) {
  const fetch = (await axios.get(
    `http://localhost:8081/api/${curr}?${prev}=${prevId}`
  )).data;
  return Object.values(await fetch)[0];
}
