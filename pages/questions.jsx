import React from "react";
import Display from "../components/Display";
import QuestionBtn from "../components/buttons/QuestionBtn";

const Questions = () => {
  return (
    <Display
      title="Questions"
      headers={["Number", "Update"]}
      items={[{ id: 2, num: 1, select: <QuestionBtn /> }]}
      buttons={[]}
    />
  );
};

export default Questions;
