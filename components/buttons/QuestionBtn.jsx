import React from "react";
import { Button } from "react-bootstrap";

const handleClick = () => {
  console.log("Updating!");
};

const QuestionBtn = props => {
  return <Button onClick={handleClick}>Update</Button>;
};

export default QuestionBtn;
