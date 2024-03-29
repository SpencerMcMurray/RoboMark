import React from "react";
import QuestionModal from "../components/modals/QuestionModal";
import TestModal from "./modals/TestModal";
import PreviewModal from "./modals/PreviewModal";

const CustomModal = props => {
  return <div className="col-12 pt-4 text-center">{props.children}</div>;
};

const ModalManager = props => {
  return (
    <CustomModal>
      <QuestionModal {...props.question} />
      <TestModal {...props.test} />
      <PreviewModal {...props.preview} />
    </CustomModal>
  );
};

export default ModalManager;
