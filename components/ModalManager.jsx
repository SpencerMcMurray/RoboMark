import React from "react";
import QuestionModal from "../components/modals/QuestionModal";

const CustomModal = props => {
  return <div className="col-12 pt-4 text-center">{props.children}</div>;
};

const ModalManager = props => {
  return (
    <CustomModal>
      <QuestionModal {...props.question} />
    </CustomModal>
  );
};

export default ModalManager;
