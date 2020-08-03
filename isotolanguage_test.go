package isotolanguage_test

import (
	"github.com/stretchr/testify/require"
	"isotolanguage"
	"testing"
)

func TestService_FindFromIso(t *testing.T) {
	err, service := isotolanguage.New()

	require.NoError(t, err)

	require.Equal(t, "English", service.FindFromIso("en").EnglishName)
	require.Equal(t, "French", service.FindFromIso("fr").EnglishName)
	require.Equal(t, "Arabic", service.FindFromIso("ar").EnglishName)
	require.Equal(t, "Chinese", service.FindFromIso("zh").EnglishName)
}