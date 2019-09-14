import React, { useState } from "react";
import Display from "../components/Display";
import QuestionBtn from "../components/buttons/QuestionBtn";
import { Button } from "react-bootstrap";

const btns = [
  <Button
    variant="success"
    onClick={() => {
      console.log("Created new");
    }}
  >
    Create New
  </Button>
];

const Questions = () => {
  const [showQuestion, setShowQuestion] = useState(false);
  return (
    <Display
      show={showQuestion}
      onHide={() => setShowQuestion(false)}
      title="Questions"
      headers={["Number", "Update"]}
      items={[
        {
          id: 2,
          num: 1,
          select: <QuestionBtn onClick={() => setShowQuestion(true)} />
        }
      ]}
      buttons={btns}
    />
  );
};

export default Questions;
