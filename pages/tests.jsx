import React from "react";
import Display from "../components/Display";
import SelectBtn from "../components/buttons/SelectBtn";
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

const Tests = () => {
  return (
    <Display
      title="Tests"
      headers={["Name", "Select"]}
      items={[{ id: 2, name: "CSC263 Final", select: <SelectBtn /> }]}
      buttons={btns}
    />
  );
};

export default Tests;
