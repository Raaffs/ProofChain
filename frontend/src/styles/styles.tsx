export const textFieldSx = {
    padding: "5px",
    
    // Target the label when focused
    '& .MuiInputLabel-root.Mui-focused': {
        color: '#E94057', // Deep Coral
    },
    
    // Target the focused underline bar
    '& .MuiInput-underline:after': {
        borderBottomColor: '#E94057', // Deep Coral
        borderBottomWidth: '2px', 
    },
    
    // Target the underline bar when hovered, before focus
    '& .MuiInput-underline:hover:not(.Mui-disabled):before': {
        borderBottomColor: '#F27121', // Lighter Orange
    },
};

export const btnstyle = {
    background: "linear-gradient(90deg, #E94057 10%, #F27121 90%)",
    padding: "8px 24px",
    color: "white",
    fontSize: "1rem",
    fontWeight: 600,
    borderRadius: "100px",
    boxShadow: "0 4px 10px 0 rgba(233, 64, 87, 0.4)",
    transition: "transform 0.2s ease-in-out",
  };

export const menuItemStyle = {
  backgroundColor: "white",
  color: "#333",
  fontWeight: 500,
  fontSize: "1rem",
  transition: "0.15s ease",

  "&:hover": {
    backgroundColor: "rgba(233, 64, 87, 0.15)", // soft red tint
    color: "#E94057",
  },

  "&.Mui-selected": {
    backgroundColor: "#E94057", // strong red for selected
    color: "white",
  },

  "&.Mui-selected:hover": {
    backgroundColor: "#D7374F", // slightly darker red on hover
  },
};
