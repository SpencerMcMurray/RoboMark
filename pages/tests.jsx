import React from "react";
import Display from "../components/Display";
import SelectBtn from "../components/buttons/SelectBtn";
import TestBtns from "../components/buttons/TestBtns";

const Tests = () => {
  return (
    <Display
      title="Tests"
      headers={["Name", "Select"]}
      items={[{ id: 2, name: "CSC263 Final", select: <SelectBtn /> }]}
      buttons={<TestBtns />}
    />
  );
};

export default Tests;
