export const naiveThemeOverrides = {
    common: {
        primaryColor: "#2080f0",
        primaryColorHover: "#4098fc",
        primaryColorPressed: "#1060c9",
        primaryColorSuppl: "#4098fc",

        borderRadius: "6px",
        borderRadiusSmall: "4px",
        borderRadiusMedium: "6px",
        borderRadiusLarge: "8px",

        fontSize: "14px",
        fontSizeSmall: "12px",
        fontSizeMedium: "14px",
        fontSizeLarge: "16px",

        heightSmall: "28px",
        heightMedium: "34px",
        heightLarge: "40px"
    }
}

export function selectTheme(dark) {
    return dark ? darkTheme : lightTheme
}

const lightTheme = {
    "common": {
        "primaryColor": "#18a058", // 主色 - 青绿色
        "primaryColorHover": "#36ad6a",
        "primaryColorPressed": "#0c7a43",
        "primaryColorSuppl": "#36ad6a",

        "infoColor": "#2080f0",   // 信息色 - 蓝色
        "infoColorHover": "#4098fc",
        "infoColorPressed": "#1060c9",

        "successColor": "#18a058", // 成功色 - 青绿色
        "successColorHover": "#36ad6a",
        "successColorPressed": "#0c7a43",

        "warningColor": "#f0a020", // 警告色 - 橙色
        "warningColorHover": "#fcb040",
        "warningColorPressed": "#c97c10",

        "errorColor": "#d03050",   // 错误色 - 红色
        "errorColorHover": "#de576d",
        "errorColorPressed": "#ab1f3f",

        "textColorBase": "#1f2937",  // 基础文本色
        "textColor1": "rgb(31, 41, 55)",
        "textColor2": "rgb(55, 65, 81)",
        "textColor3": "rgb(75, 85, 99)",

        "borderRadius": "6px",
        "borderRadiusSmall": "4px",
        "borderRadiusMedium": "6px",
        "borderRadiusLarge": "8px",

        "fontSize": "14px",
        "fontSizeSmall": "12px",
        "fontSizeMedium": "14px",
        "fontSizeLarge": "16px",

        "heightSmall": "28px",
        "heightMedium": "34px",
        "heightLarge": "40px"
    },
    "Button": {
        "colorPrimary": "#18a058",
        "textColorPrimary": "#ffffff",
        "borderPrimary": "#18a058",
        "borderRadiusMedium": "6px",
        "borderRadiusLarge": "8px"
    },
    "Input": {
        "borderRadius": "6px",
        "borderFocus": "#18a058",
        "borderHover": "#36ad6a",
        "caretColor": "#18a058"
    },
    "Card": {
        "borderRadius": "8px",
        "color": "#ffffff",
        "colorModal": "#ffffff"
    },
    "Menu": {
        "itemColorActive": "rgba(24, 160, 88, 0.1)",
        "itemTextColorActive": "#18a058",
        "itemIconColorActive": "#18a058"
    },
    "Tag": {
        "borderRadius": "4px"
    },
    "Alert": {
        "borderRadius": "6px"
    }
}

const darkTheme = {
    "common": {
        "primaryColor": "#63e2b7", // 更亮的青绿色适应暗色背景
        "primaryColorHover": "#7fe7c4",
        "primaryColorPressed": "#5acea7",
        "primaryColorSuppl": "#7fe7c4",

        "infoColor": "#70c0e8",
        "infoColorHover": "#8acbec",
        "infoColorPressed": "#66afd3",

        "successColor": "#63e2b7",
        "successColorHover": "#7fe7c4",
        "successColorPressed": "#5acea7",

        "warningColor": "#f2c97d",
        "warningColorHover": "#f5d599",
        "warningColorPressed": "#e6c260",

        "errorColor": "#e88080",
        "errorColorHover": "#e98b8b",
        "errorColorPressed": "#e57272",

        // 暗色主题文本色
        "textColorBase": "#e5e7eb",
        "textColor1": "rgb(229, 231, 235)",
        "textColor2": "rgb(209, 213, 219)",
        "textColor3": "rgb(156, 163, 175)",

        // 背景色
        "bodyColor": "#1a1a1a",
        "cardColor": "#242424",
        "modalColor": "#242424",
        "popoverColor": "#242424",

        "borderRadius": "6px",
        "borderRadiusSmall": "4px",
        "borderRadiusMedium": "6px",
        "borderRadiusLarge": "8px"
    },
    "Button": {
        "colorPrimary": "#63e2b7",
        "textColorPrimary": "#1a1a1a"
    },
    "Card": {
        "color": "#242424",
        "colorModal": "#2a2a2a"
    },
    "Menu": {
        "itemColorActive": "rgba(99, 226, 183, 0.2)",
        "itemTextColorActive": "#63e2b7"
    }
}
