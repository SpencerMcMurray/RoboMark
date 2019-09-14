import React from "react";
import Display from "../components/Display";
import SelectBtn from "../components/buttons/SelectBtn";

const Tests = () => {
  return (
    <Display
      title="Tests"
      headers={["Name", "Select"]}
      items={[{ id: 2, name: "CSC263 Final", select: <SelectBtn /> }]}
      buttons={[]}
    />
  );
};

export default Tests;
