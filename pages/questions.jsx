import React from "react";
import Display from "../components/Display";
import SelectBtn from "../components/buttons/SelectBtn";

const Questions = () => {
  return (
    <Display
      title="Questions"
      headers={["Number", "Update"]}
      items={[{ id: 2, num: 1, select: <SelectBtn /> }]}
      buttons={[]}
    />
  );
};

export default Questions;
