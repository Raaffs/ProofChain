import React from "react";
import { Dialog, DialogContent } from "@mui/material";
import IssueCard, { CertificateData } from "../cards/certificate";

interface ModalIssueCardProps {
  open: boolean;
  onClose: () => void;
  data: CertificateData | null;
  viewTitle: string;
  onIssue: (cert: CertificateData) => void;
}

const ModalIssueCard: React.FC<ModalIssueCardProps> = ({
  open,
  onClose,
  data,
  viewTitle,
  onIssue,
}) => {
  return (
    <Dialog
      open={open}
      onClose={onClose}
      maxWidth="md"
      fullWidth={false} // 👈 important — otherwise MUI forces centering
      sx={{
        "& .MuiDialog-paper": {
          position: "fixed",
          right: 0,
          top: 0,
          height: "100vh",
          margin: 0,
          borderRadius: 0, // optional
        },
      }}
    >
      <DialogContent sx={{ p: 0 }}>
        <IssueCard data={data} viewTitle={viewTitle} onIssue={onIssue} />
      </DialogContent>
    </Dialog>
  );
};

export default ModalIssueCard;
