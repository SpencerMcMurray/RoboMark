import React, { useState } from "react";
import Display from "../components/Display";
import ScanBtn from "../components/buttons/ScanBtn";
import { Button } from "react-bootstrap";

const Pages = () => {
  const [show, setShow] = useState(false);

  const btns = [
    <Button
      variant="success"
      onClick={() => {
        setShow(true);
      }}
    >
      Create New
    </Button>,
    <Button
      onClick={() => {
        console.log("Finshed");
      }}
    >
      Finish
    </Button>
  ];

  return (
    <Display
      showPage={show}
      onHidePage={() => setShow(false)}
      title="Pages"
      headers={["Number", "Scan"]}
      items={[{ id: 2, name: 1, scan: <ScanBtn id={2} /> }]}
      buttons={btns}
    />
  );
};

export default Pages;
