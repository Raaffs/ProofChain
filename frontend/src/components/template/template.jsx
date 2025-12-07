import { Card, Typography } from "@mui/material";
import {useTheme} from '@emotion/react'
import { tokens } from "../../themes";
export function PlatformOverviewCard() {
  const theme=useTheme()
  const colors=tokens(theme.palette.mode)
  return (
    <Card
      sx={{
        padding: "10px",
        width: "100%",
        minHeight: "60%",
        borderRadius: "16px",
        marginLeft: "20px",
        backgroundRepeat: "no-repeat",
        backgroundColor:
          theme.palette.mode === "dark" ? "transparent" : "white",
      }}
    >
      <Typography
        variant="h2"
        align="center"
        sx={{
          background:
            "linear-gradient(90deg, #ff4d4d 0%, #ff6f61 40%, #ff8c75 80%, #ffc1a1 100%)",
          WebkitBackgroundClip: "text",
          WebkitTextFillColor: "transparent",
          fontFamily: "'Poppins', sans-serif",
          mb: 1,
          fontWeight: 700,
          textShadow: "0px 0px 10px rgba(255, 120, 120, 0.35)",
          borderBottom: "2px solid #ff4d4d",
          display: "inline-block",
          paddingBottom: "4px",
          width: "100%",
        }}
      >
        ProofChain
      </Typography>

      <Typography
        variant="body1"
        align="left"
        sx={{
          mt: 2,
          fontSize: "1.1rem",
          lineHeight: 1.6,
          color: colors.grey[100] || "#fafafa",
          padding: "0 12px",
        }}
      >
        Our platform lets you share only the information you choose while
        still proving your documents are authentic. Institutions issue secure
        digital records, and you decide exactly which details to reveal.
        Everything is protected with strong encryption, stored safely, and
        verifiable on the spot — without exposing anything you want to keep
        private. It’s a simple, safe way to issue, store, and verify important
        documents.
      </Typography>
    </Card>
  );
}
