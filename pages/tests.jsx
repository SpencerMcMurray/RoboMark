import React, { useState } from "react";
import Display from "../components/Display";
import { Button } from "react-bootstrap";
import LoadingScreen from "../components/LoadingScreen";

import fetchData from "../static/helpers/fetchData";
import { useRouter } from "next/router";
import { useAsync, IfPending, IfFulfilled } from "react-async";
const axios = require("axios");

const Tests = () => {
  const [show, setShow] = useState(false);
  const router = useRouter();

  const btns = [
    <Button
      onClick={() => {
        setShow(true);
      }}
    >
      Create New
    </Button>
  ];

  const allTests = useAsync({
    promiseFn: fetchData,
    curr: "tests",
    prev: "users",
    prevId: router.query.user
  });

  const addTest = async item => {
    const postData = {
      user: router.query.user,
      name: item.name,
      markDenom: 0
    };
    axios
      .post("http://localhost:8081/api/tests", postData)
      .then(() => allTests.reload())
      .then(() => setShow(false));
  };

  return (
    <React.Fragment>
      <IfPending state={allTests}>
        <LoadingScreen />
      </IfPending>
      <IfFulfilled state={allTests}>
        <Display
          showTest={show}
          onHideTest={() => setShow(false)}
          title="Tests"
          headers={["Name", "Select"]}
          items={allTests.value}
          buttons={btns}
          add={addTest}
        />
      </IfFulfilled>
    </React.Fragment>
  );
};

export default Tests;
