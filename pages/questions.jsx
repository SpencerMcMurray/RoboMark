import React, { useState } from "react";
import Display from "../components/Display";
import QuestionBtn from "../components/buttons/QuestionBtn";
import { Button } from "react-bootstrap";

const Questions = () => {
  const [showPreview, setShowPreview] = useState(false);
  const [show, setShow] = useState(false);

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

  return (
    <Display
      showPreview={showPreview}
      onHidePreview={() => setShowPreview(false)}
      showQuestion={show}
      onHideQuestion={() => setShow(false)}
      title="Questions"
      headers={["Number", "Update"]}
      items={[
        {
          id: 2,
          num: 1,
          select: <QuestionBtn onClick={() => setShowPreview(true)} />
        }
      ]}
      buttons={btns}
    />
  );
};

export default Questions;
