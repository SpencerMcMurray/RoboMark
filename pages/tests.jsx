import React, { useState } from "react";
import Display from "../components/Display";
import SelectBtn from "../components/buttons/SelectBtn";
import { Button } from "react-bootstrap";

const Tests = () => {
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
      showTest={show}
      onHideTest={() => setShow(false)}
      title="Tests"
      headers={["Name", "Select"]}
      items={[{ id: 2, name: "CSC263 Final", select: <SelectBtn /> }]}
      buttons={btns}
    />
  );
};

export default Tests;
