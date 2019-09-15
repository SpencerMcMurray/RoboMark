import React, { useState } from "react";
import Display from "../components/Display";
import QuestionBtn from "../components/buttons/QuestionBtn";
import { Button } from "react-bootstrap";
import LoadingScreen from "../components/LoadingScreen";

import fetchData from "../static/helpers/fetchData";
import { useRouter } from "next/router";
import { useAsync, IfPending, IfFulfilled } from "react-async";
const axios = require("axios");

const Questions = () => {
  const [showPreview, setShowPreview] = useState(false);
  const [show, setShow] = useState(false);

  const router = useRouter();

  const allQuestions = useAsync({
    promiseFn: fetchData,
    curr: "questions",
    prev: "tests",
    prevId: router.query.test
  });

  const btns = [
    <Button
      variant="success"
      onClick={() => {
        setShow(true);
      }}
    >
      Create New
    </Button>
  ];

  const addQuestion = async item => {
    const fetch = (await axios.get(
      `http://localhost:8081/api/tests?id=${router.query.test}`
    )).data.tests[0];
    const postData = {
      test: router.query.test,
      question: item.question,
      answers: item.answers
    };
    postData.user = (await fetch).user;
    axios
      .post("http://localhost:8081/api/questions", postData)
      .then(() => allQuestions.reload())
      .then(() => setShow(false));
  };

  return (
    <React.Fragment>
      <IfPending state={allQuestions}>
        <LoadingScreen />
      </IfPending>
      <IfFulfilled state={allQuestions}>
        <Display
          showPreview={showPreview}
          onHidePreview={() => setShowPreview(false)}
          onShowPreview={() => setShowPreview(true)}
          showQuestion={show}
          onHideQuestion={() => setShow(false)}
          title="Questions"
          headers={["Number", "Preview"]}
          items={allQuestions.value}
          buttons={btns}
          add={addQuestion}
        />
      </IfFulfilled>
    </React.Fragment>
  );
};

export default Questions;
